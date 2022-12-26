resource "aws_subnet" "private-1a" {
  vpc_id            = aws_vpc.workshop.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "ap-southeast-1a"

  tags = {
    Cluster = "training"
  }
}

resource "aws_subnet" "private-1b" {
  vpc_id            = aws_vpc.workshop.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "ap-southeast-1b"

  tags = {
    Cluster = "training"
  }
}

resource "aws_subnet" "public-1a" {
  vpc_id                  = aws_vpc.workshop.id
  cidr_block              = "10.0.3.0/24"
  availability_zone       = "ap-southeast-1a"
  map_public_ip_on_launch = true

  tags = {
    Cluster = "training"
  }
}
resource "aws_subnet" "public-1b" {
  vpc_id                  = aws_vpc.workshop.id
  cidr_block              = "10.0.4.0/24"
  availability_zone       = "ap-southeast-1b"
  map_public_ip_on_launch = true

  tags = {
    Cluster = "training"
  }
}