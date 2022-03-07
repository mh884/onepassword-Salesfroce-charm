Connected onepassword with salesforce login
with charm interface we can interact with onepassword item to 
- login 
- filter
- get session

With a Secrets Automation workflow, you can securely access your 1Password items and vaults in your company’s apps and cloud infrastructure using a private REST API provided by a 1Password Connect server and login to salesforce org
Before you get started, you’ll need a deployment environment with Docker or Kubernetes to deploy the Connect server.
https://support.1password.com/secrets-automation/

Create .env file 
- OP_CONNECT_TOKEN // onepassword token
- OP_CONNECT_HOST  // onepassword docker hosted utl
- VAULT // onepassword valut

