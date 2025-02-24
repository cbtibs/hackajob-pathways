# AWS Challenge 5
```bash
# Set up Keys with least privliege access outside of the terminal
aws configure

aws ec2 create-key-pair --key-name hackajob --query 'KeyMaterial' --output text > ~/Hackajob.pem
chmod 600 ~/Hackajob.pem

aws ec2 describe-images --owners amazon --filters "Name=name,Values=amzn2-ami-hvm-*-x86_64-gp2" --query 'Images | sort_by(@, &CreationDate) | [-1].ImageId' --output text

aws ec2 run-instances --image-id ami-08a28be5eae6c1d68 --count 1 --instance-type t2.micro --key-name hackajob --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=hackajob-ec2-1}]'

aws ec2 run-instances --image-id ami-08a28be5eae6c1d68 --count 1 --instance-type t2.micro --key-name hackajob --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=hackajob-ec2-2}]'

# Default SG and Default subnet
aws elb create-load-balancer --load-balancer-name hackajob-lb \
    --listeners "Protocol=HTTP,LoadBalancerPort=80,InstanceProtocol=HTTP,InstancePort=80" \
    --subnets subnet-0630e323d932fe4e3 \ 
    --security-groups sg-0075df96fe4bb4824

aws elb register-instances-with-load-balancer --load-balancer-name hackajob-lb \
    --instances i-0f2b4d7e8c9a1b2c3 i-7e6a5b4c3d2e1f0a9
```