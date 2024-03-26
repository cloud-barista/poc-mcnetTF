
resource "aws_instance" "test_ec2_instance" {
  ami           = "ami-0f3a440bbcff3d043" # Ubuntu Server 22.04 LTS (HVM), SSD Volume Type
  instance_type = "t2.micro"

  vpc_security_group_ids = [aws_security_group.test_sg.id]

  subnet_id = aws_subnet.test_subnet_1.id
  
  tags = {
    Name = "tofu-aws-ec2-instance"
  }
}