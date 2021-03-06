package cmd

import (
	"tubectl/api"

	"github.com/spf13/cobra"
)

var image string
var replica int32
var deployName string

func init() {
	//fmt.Println("in deployment")
	createCmd.AddCommand(createDeploymentCmd)
	deleteCmd.AddCommand(deleteDeploymentCmd)
	getCmd.AddCommand(getDeploymentCmd)
	updateCmd.AddCommand(updateDeploymentCmd)

	updateCmd.PersistentFlags().StringVarP(&deployName, "deploy", "d", "demo-deployment", "This flag is to give the deployment name which need to be updated")
	updateCmd.PersistentFlags().StringVarP(&image, "image", "i", "nginx:1.12", "This flag is to give the image name for updateing the deployment image")
	updateCmd.PersistentFlags().Int32VarP(&replica, "replica", "r", 2, "This flag is to give the replica number for updating the replica count of the deployment")
}

//createDeploymentCmd is a variable
var createDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of create command, it is used to ",
	Long:  "This command is for creating a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.CreateDeployment()
	},
}

//deleteDeploymentCmd is a variable
var deleteDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of delete command, it is used to ",
	Long:  "This command is for deleting a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.DeleteDeployment(args)
	},
}

//getDeploymentCmd is a variable
var getDeploymentCmd = &cobra.Command{
	Use:   "deployments",
	Short: "deployment is a sub-command of get command, it is used to get all the deployment in default namespace",
	Long:  "This command is for getting all the deployment object in default namespace using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.GetDeployments()
	},
}

//updateDeploymentCmd is a variable
var updateDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of create command, it is used to ",
	Long:  "This command is for creating a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.UpdateDeployment(deployName, replica, image)
	},
}
