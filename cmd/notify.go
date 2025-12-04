package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bazgab/opencpe/policies"
	"github.com/bazgab/opencpe/utils/errors"
	"github.com/bazgab/opencpe/utils/logging"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"log/slog"
	"net"
	"net/smtp"
	"os"
	"time"
)

type Config struct {
	Authentication struct {
		AwsProfile     string `json:"aws_profile"`
		AwsAccountId   int    `json:"aws_account_id"`
		AwsAccountName string `json:"aws_account_name"`
	} `json:"authentication"`

	Notification struct {
		SmtpEndpoint string `json:"smtp_endpoint"`
		SmtpPort     int    `json:"smtp_port"`
		SmtpUser     string `json:"smtp_user"`
		SmtpPassword string `json:"smtp_password"`
		SenderEmail  string `json:"sender_email"`
	} `json:"notification"`

	IgnoredTags struct {
		Owner   []string `json:"owner"`
		Project []string `json:"project"`
	} `json:"ignored_tags"`
}

type EmailData struct {
	Policy       string
	InstanceName string
	InstanceId   string
	Region       string
	AwsAccount   string
}

var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "notify resource owners of policy infringement",
	Long:  `This command only notifies resource owners of policy infringement, as opposed to notify-and-delete.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Check if the global flag "policy" was actually set by the user
		RequiredFlags := []string{"config", "policy", "region"}
		for _, flag := range RequiredFlags {
			if !cmd.Flags().Changed(flag) {
				return fmt.Errorf("required flag %s not set for the 'notify' command", flag)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Checking if global flags are working
		fmt.Println("OpenCPE - Notify")

		cfgFile, err := os.ReadFile(flagConfig)
		if err != nil {
			slog.Error("Error when opening file: ", err)
		}

		var cfg Config

		err = json.Unmarshal(cfgFile, &cfg)
		if err != nil {
			slog.Error("Error during Unmarshal(): ", err)
		}
		logging.BreakerLine()
		fmt.Println("")
		fmt.Println("Loaded Configuration:")
		fmt.Printf("-- Authentication.Profile: %s\n", cfg.Authentication.AwsProfile)
		fmt.Printf("-- Authentication.Account_Id: %d\n", cfg.Authentication.AwsAccountId)
		fmt.Printf("-- Authentication.Account_Name: %s\n", cfg.Authentication.AwsAccountName)
		fmt.Printf("-- Notification.SMTP_Endpoint: %s\n", cfg.Notification.SmtpEndpoint)
		fmt.Printf("-- Notification.SMTP_Port: %d\n", cfg.Notification.SmtpPort)
		fmt.Printf("-- Notification.SMTP_User: %s\n", cfg.Notification.SmtpUser)
		fmt.Printf("-- Notification.Sender_Email: %s\n", cfg.Notification.SenderEmail)
		for _, owner := range cfg.IgnoredTags.Owner {
			fmt.Printf("-- IgnoredTags.Owner: %s\n", owner)
		}
		for _, project := range cfg.IgnoredTags.Project {
			fmt.Printf("-- IgnoredTags.Project: %s\n", project)
		}
		logging.BreakerLine()
		fmt.Println()

		errors.IdentityCheck(cfg.Authentication.AwsProfile, flagRegion, cfg.Authentication.AwsAccountId)

		//Check for policy
		if flagPolicy == "instance-age-2-days" {
			fmt.Println("")
			fmt.Println("Policy: instance-age-2-days")
			fmt.Printf("Profile: %s\n", cfg.Authentication.AwsProfile)
			fmt.Printf("Region: %s\n", flagRegion)

			appConfig, _ := policies.LoadConfig(flagConfig)

			instances, err := policies.InstanceAge2Days(cfg.Authentication.AwsProfile, flagRegion, appConfig.IgnoredTags)
			if err != nil {
				log.Fatal(err)
			}

			logging.BreakerLine()
			fmt.Println("")
			fmt.Printf("Received %d instances.\n", len(instances))

			// Parsing files before the loop for memory efficiency
			t, err := template.ParseFiles("utils/templates/email_template.html")
			if err != nil {
				log.Fatal("Could not parse template:", err)
			}
			// INSIDE YOUR LOOP
			for _, inst := range instances {
				fmt.Printf("[Instance Name: %s | Instance Id: %s]\n", inst.Name, inst.Id)

				// --- 1. PREPARE CONTENT ---
				d := EmailData{
					InstanceName: inst.Name,
					InstanceId:   inst.Id,
					Region:       flagRegion,
					AwsAccount:   cfg.Authentication.AwsAccountName,
				}

				var body bytes.Buffer
				if err := t.Execute(&body, d); err != nil {
					log.Printf("❌ Template error: %v", err)
					continue
				}

				// Header Fix

				headers := "From: " + cfg.Notification.SenderEmail + "\r\n" +
					"To: " + inst.Owner + "\r\n" +
					"Subject: Resource Action Required on Instance " + inst.Name + "!\r\n" +
					"MIME-Version: 1.0\r\n" +
					"Content-Type: text/html; charset=UTF-8\r\n" +
					"\r\n"

				msg := []byte(headers + body.String())

				// --- 2. CONNECTION SETUP (PORT 587 SPECIFIC) ---
				addr := fmt.Sprintf("%s:%d", cfg.Notification.SmtpEndpoint, cfg.Notification.SmtpPort)

				fmt.Printf("⏳ Attempting connection to %s with 5s timeout...\n", addr)

				// Use DialTimeout. If this errors, your Network is blocking you.
				conn, err := net.DialTimeout("tcp4", addr, 5*time.Second)
				if err != nil {
					log.Printf("❌ NETWORK ERROR: Your firewall/ISP/Cloud Provider is blocking Port 587.\nError details: %v", err)
					continue
				}

				fmt.Println("✅ TCP Connection established! (The issue was not the network)")
				// A. Plain TCP Connection (Force IPv4)
				// We use net.Dial, NOT tls.Dial because 587 starts as plain text
				fmt.Println("Dialing Port 587 via IPv4...")
				conn, err = net.Dial("tcp4", addr)
				if err != nil {
					log.Printf("❌ Connection failed: %v", err)
					continue
				}

				// B. Create SMTP Client
				c, err := smtp.NewClient(conn, cfg.Notification.SmtpEndpoint)
				if err != nil {
					log.Printf("❌ Client creation failed: %v", err)
					conn.Close() // Close raw connection if client fails
					continue
				}

				// C. Upgrade to TLS (StartTLS) - REQUIRED for Port 587
				tlsConfig := &tls.Config{
					InsecureSkipVerify: false,
					ServerName:         cfg.Notification.SmtpEndpoint,
				}
				if err = c.StartTLS(tlsConfig); err != nil {
					log.Printf("❌ StartTLS failed: %v", err)
					c.Quit()
					continue
				}

				// --- 3. AUTHENTICATE ---
				// Must happen AFTER StartTLS
				auth := smtp.PlainAuth("", cfg.Notification.SmtpUser, cfg.Notification.SmtpPassword, cfg.Notification.SmtpEndpoint)
				if err = c.Auth(auth); err != nil {
					log.Printf("❌ Authentication failed: %v", err)
					c.Quit()
					continue
				}

				// --- 4. SEND EMAIL ---
				if err = c.Mail(cfg.Notification.SenderEmail); err != nil {
					log.Printf("❌ Sending failure on SenderEmail: %v", err)
					c.Quit()
					continue
				}

				if err = c.Rcpt(inst.Owner); err != nil {
					log.Printf("❌ Sending failure on Recipient side: %v", err)
					c.Quit()
					continue
				}

				w, err := c.Data()
				if err != nil {
					log.Printf("❌ Data command failed: %v", err)
					c.Quit()
					continue
				}

				_, err = w.Write(msg)
				if err != nil {
					log.Printf("❌ Write failed: %v", err)
					c.Quit()
					continue
				}

				err = w.Close()
				if err != nil {
					log.Printf("❌ Close Data failed: %v", err)
					c.Quit()
					continue
				}

				// --- 5. CLEANUP ---
				c.Quit()
				fmt.Println("✅ Success! Email sent via Port 587.")
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)

}
