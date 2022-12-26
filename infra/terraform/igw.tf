resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.workshop.id

  tags = {
    Cluster = "training"
  }
}