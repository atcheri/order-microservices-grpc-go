# Tiltfile for order microservices

# Load environment variables from .env file if it exists
load('ext://dotenv', 'dotenv')
dotenv()

# MySQL Database
# Read the Kubernetes YAML and substitute the password
mysql_yaml = str(read_file('infra/mysql.yaml'))
mysql_yaml = mysql_yaml.replace('MYSQL_ROOT_PASSWORD_PLACEHOLDER', os.getenv('MYSQL_ROOT_PASSWORD', 'defaultpassword'))
k8s_yaml(blob(mysql_yaml))

# Port forward to access MySQL locally
k8s_resource('mysql', port_forwards='3306:3306')
