package main

import (
	"context"
	"net/http"

	dashboard "github.com/2023mt03014/unused-cloud-resources/handler"

	"github.com/gin-gonic/gin"
	aws_unused "github.com/2023mt03014/unused-cloud-resources/aws_unused_resources"
	gcp_unused "github.com/2023mt03014/unused-cloud-resources/gcp_unused_resources"
)

func getAWSEBSData(c *gin.Context) {
	unused_ebs_data := aws_unused.Get_unused_ebs_volumes("us-east-1")
	percentage_unused := 100 * unused_ebs_data.UnusedInstancesCount / unused_ebs_data.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unused_ebs_data.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unused_ebs_data.TotalInstancesCount,
		"unused_count": unused_ebs_data.UnusedInstancesCount,
	})
}

func getAWSEC2Data(c *gin.Context) {
	unused_ec2_data := aws_unused.GetUnusedEC2Instances(context.Background(),
		"us-east-1",
		1.0,
		7)
	percentage_unused := 100 * unused_ec2_data.UnusedInstancesCount / unused_ec2_data.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unused_ec2_data.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unused_ec2_data.TotalInstancesCount,
		"unused_count": unused_ec2_data.UnusedInstancesCount,
	})
}

func getAWSS3Data(c *gin.Context) {
	unused_s3_data, _ := aws_unused.GetUnusedS3Buckets(context.Background(), "us-east-1", 1, 7)
	percentage_unused := 100 * unused_s3_data.UnusedInstancesCount / unused_s3_data.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unused_s3_data.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unused_s3_data.TotalInstancesCount,
		"unused_count": unused_s3_data.UnusedInstancesCount,
	})
}

func getAWSVPCData(c *gin.Context) {
	unused_vpcs_data, _ := aws_unused.GetUnusedVPCs(context.Background(), "us-east-1", 0)
	percentage_unused := 100 * unused_vpcs_data.UnusedInstancesCount / unused_vpcs_data.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unused_vpcs_data.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unused_vpcs_data.TotalInstancesCount,
		"unused_count": unused_vpcs_data.UnusedInstancesCount,
	})
}

func getAWSRDSData(c *gin.Context) {
	unused_rds_data, _ := aws_unused.GetUnusedRDSInstances(context.Background(), "us-east-1", 5.0, 7)
	percentage_unused := 100 * unused_rds_data.UnusedInstancesCount / unused_rds_data.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unused_rds_data.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unused_rds_data.TotalInstancesCount,
		"unused_count": unused_rds_data.UnusedInstancesCount,
	})
}

func getAWSLambdaData(c *gin.Context) {
	unused_lambda_data, _ := aws_unused.GetUnusedLoadBalancers(context.Background(), "us-east-1", 100, 7)
	percentage_unused := 100 * unused_lambda_data.UnusedInstancesCount / unused_lambda_data.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unused_lambda_data.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unused_lambda_data.TotalInstancesCount,
		"unused_count": unused_lambda_data.UnusedInstancesCount,
	})
}

func getGCPUnusedDisk(c *gin.Context) {
	unused_disk_data := gcp_unused.Get_Unused_Disks("finops-accelerator", "us-central1-a")
	percentage_unused := 100 * unused_disk_data.UnusedInstancesCount / unused_disk_data.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unused_disk_data.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unused_disk_data.TotalInstancesCount,
		"unused_count": unused_disk_data.UnusedInstancesCount,
	})
}

func getGCPUnusedIP(c *gin.Context) {
	unusedIPsData := gcp_unused.Get_Unused_IPs("finops-accelerator", "us-central1")
	percentage_unused := 100 * unusedIPsData.UnusedInstancesCount / unusedIPsData.TotalInstancesCount
	c.JSON(http.StatusOK, gin.H{
		"resource_ids": unusedIPsData.ResourceIDs,
		"percentage":   percentage_unused,
		"total_count":  unusedIPsData.TotalInstancesCount,
		"unused_count": unusedIPsData.UnusedInstancesCount,
	})
}

func main() {
	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Okay",
		})
	})

	r.GET("/aws/ebs", getAWSEBSData)
	r.GET("/aws/ec2", getAWSEC2Data)
	r.GET("/aws/s3", getAWSS3Data)
	r.GET("/aws/vpc", getAWSVPCData)
	r.GET("/aws/rds", getAWSRDSData)
	r.GET("/aws/lambda", getAWSLambdaData)
	r.GET("/gcp/disks", getGCPUnusedDisk)
	r.GET("/gcp/ips", getGCPUnusedIP)
	r.GET("/dashboard", dashboard.DashboardHandler)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
