The project aims to find a fast way to log in to salesforce for the developer and our app in the pipeline to sync with the QA regression test.
Connected 1password with salesforce login with charm interface to interact with 1password item to 
- login
- filter
- -get a session inside the command line

With a Secrets Automation workflow, you can securely access your 1Password items and vaults in your company’s apps and cloud infrastructure using a private REST API provided by a 1Password Connect server and login to salesforce org. Before you get started, you’ll need a deployment environment with Docker or Kubernetes to deploy the Connect.

https://support.1password.com/secrets-automation/

Create .env file 
- OP_CONNECT_TOKEN // onepassword token
- OP_CONNECT_HOST  // onepassword docker hosted utl
- VAULT // onepassword valut

