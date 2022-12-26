resource "aws_iam_role" "eks_iam" {
  name = "eks-cluster-demo"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "eks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "eks_iam-AmazonEKSClusterPolicy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.eks_iam.name
}

resource "aws_eks_cluster" "devops-cluster" {
  name     = "eks-devops-cluster"
  role_arn = aws_iam_role.eks_iam.arn

  vpc_config {
    subnet_ids = [
      aws_subnet.private-1a.id,
      aws_subnet.private-1b.id,
      aws_subnet.public-1a.id,
      aws_subnet.public-1b.id
    ]
  }

  depends_on = [aws_iam_role_policy_attachment.eks_iam-AmazonEKSClusterPolicy]
}