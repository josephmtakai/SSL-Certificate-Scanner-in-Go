# SSL Certificate Scanner

Welcome to the **SSL Certificate Scanner**! This Go application allows users to scan websites for SSL certificate details, including expiration status, issuer information, and cipher suite validation.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)
- [Contributing](#contributing)
- [Contact](#contact)

## Features

- **SSL Expiration Checker**: Check if SSL certificates are expired or valid.
- **Issuer Information**: View details about the certificate issuer.
- **Cipher Suite Validation**: Identify weak cipher suites used by the website.
- **User-friendly Interface**: Simple command-line interface for easy interaction.

## Requirements

- [Go](https://golang.org/dl/) (version 1.16 or higher)

## Installation

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/yourusername/ssl-certificate-scanner.git
Navigate to the project directory:

cd ssl-certificate-scanner
Run the application:

go run main.go
Usage
Start the SSL Certificate Scanner by running the following command:


go run main.go
Enter the website URL you wish to scan (e.g., [https://www.example.com](https://aws.amazon.com/partners/training/)).

The scanner will display the following information:

Certificate Subject
Issuer
Not Before and Not After dates
Expiration Status
Cipher Suite Strength
To exit the program, type exit.

Example Output
Welcome to the SSL Certificate Scanner
Enter a website URL (or type 'exit' to quit): https://www.example.com
Scan Results:
Certificate Subject: ...
Issuer: ...
Not Before: ...
Not After: ...
Status: VALID
Cipher Suite: Strong

License
This project is licensed under the MIT License. See the LICENSE file for details.

Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

Fork the repository.
Create your feature branch:

git checkout -b feature/new-feature
Commit your changes:

git commit -m "Add a new feature"
Push to the branch:

git push origin feature/new-feature
Open a pull request.

SCREENSHOT: ![ssl](https://github.com/user-attachments/assets/f43ebbf2-d613-4a9a-8aa9-212b5f2c1cbd)


Contact
For any inquiries or suggestions, please feel free to contact me at:
Email: joseph.mtakai@outlook.com
LinkedIn: [Your LinkedIn Profile](https://www.linkedin.com/in/joseph-n-mtakai-b6b94a76/)

### Notes:
1. **Customize the Content**: Make sure to replace placeholders like `yourusername`, `your-email@example.com`, and links to your actual GitHub repository and LinkedIn profile.
2. **Additional Sections**: You can add more sections, like `Future Improvements` or `Troubleshooting`, if needed.
3. **LICENSE File**: If you're including a license, create a LICENSE file in the same directory as your README file.

This README provides a comprehensive overview of your project, making it easy for others to understand how to use it and contribute. Let me know if you need any more modifications or additional sections!
