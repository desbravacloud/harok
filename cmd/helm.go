/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package rocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	helmclient "github.com/mittwald/go-helm-client"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/client-go/tools/clientcmd"
)

// convert bytes to string
func bytesToString(data []byte) string {
	return string(data[:])
}

// runHelm run a helm command from current context
func runHelm(context, namespace string) helmclient.Client {
	home := os.Getenv("HOME")
	kubeConfigPath := filepath.Join(home, ".kube", "config")

	setK8sContext(context, kubeConfigPath)

	helm, err := helmclient.New(&helmclient.Options{
		Namespace:        namespace,
		RepositoryCache:  filepath.Join(home, ".cache/helm/repository"),
		RepositoryConfig: filepath.Join(home, "/.config/helm/repositories.yaml"),
		Linting:          true,
		Debug:            true,
		DebugLog: func(format string, v ...interface{}) {
			log.Printf(format, v)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return helm
}

//getHelmRelease get values from helm release in Source K8s Cluster
func getHelmRelease(release, context, namespace string) release.Release {
	getRelease, err := runHelm(context, namespace).GetRelease(release)
	if err != nil {
		log.Fatal(err)
	}

	return *getRelease
}

func migrateRelease(release, repo, namespace, sourceContext, targetContext string) {
	releaseConfig := getHelmRelease(release, sourceContext, namespace)

	valuesJson, err := json.Marshal(releaseConfig.Config)
	if err != nil {
		log.Fatal(err)
	}
	
	values := bytesToString(valuesJson)
	
	if _, err := runHelm(targetContext, releaseConfig.Namespace).InstallOrUpgradeChart(context.Background(), &helmclient.ChartSpec{
		ReleaseName: releaseConfig.Name,
		ChartName:   filepath.Join(repo, releaseConfig.Name),
		Namespace:   releaseConfig.Namespace,
		ValuesYaml:  values,
		UpgradeCRDs: false,
		Wait:        false,
	}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Release \"%s\" migrated to cluster \"%s\"", releaseConfig.Name, targetContext)
	}
}

// Set your K8s context to run helm commands
func setK8sContext(context, kubeconfigPath string) (err error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath}
	configOverrides := &clientcmd.ConfigOverrides{}

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.RawConfig()
	if err != nil {
		return fmt.Errorf("Error getting RawConfig: %w", err)
	}

	if config.Contexts[context] == nil {
		return fmt.Errorf("Context %s doesn't exists", context)
	}

	config.CurrentContext = context
	err = clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), config, true)
	if err != nil {
		return fmt.Errorf("Error ModifyConfig: %w", err)
	}

	//fmt.Printf("Switched to context \"%s\"", context)
	return nil
}

// helmCmd represents the helm command
var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Manage helm releases",
	Long:  `Manage your helm charts deployed in your Kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {},
}

//migrateCmd migrate applications between clusters using helm
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Helm subcommand to migrate apps",
	Long:  `Migrate applications between Kubernetes clusters using helm`,
	Run: func(cmd *cobra.Command, args []string) {

		release, err := cmd.PersistentFlags().GetString("release")
		if err != nil {
			log.Fatal(err)
		}

		repo, err := cmd.PersistentFlags().GetString("repo")
		if err != nil {
			log.Fatal(err)
		}

		namespace, err := cmd.PersistentFlags().GetString("namespace")
		if err != nil {
			log.Fatal(err)
		}

		sourceContext, err := cmd.PersistentFlags().GetString("source")
		if err != nil {
			log.Fatal(err)
		}

		targetContext, err := cmd.PersistentFlags().GetString("target")
		if err != nil {
			log.Fatal(err)
		}

		migrateRelease(release, repo, namespace, sourceContext, targetContext)	
	},
}

func init() {
	// Add commands to root command
	rootCmd.AddCommand(helmCmd)

	// Add commands to helm command
	helmCmd.AddCommand(migrateCmd)

	// Add flags to migrate subcommand
	migrateCmd.PersistentFlags().StringP("release", "r", "", "Helm release name")
	migrateCmd.MarkPersistentFlagRequired("release")
	migrateCmd.PersistentFlags().StringP("repo", "", "", "Helm repository name")
	migrateCmd.MarkPersistentFlagRequired("repo")
	migrateCmd.PersistentFlags().StringP("namespace", "n", "", "Kubernetes namespace")
	migrateCmd.MarkPersistentFlagRequired("namespace")
	migrateCmd.PersistentFlags().StringP("source", "s", "", "Source Kubernetes cluster")
	migrateCmd.MarkPersistentFlagRequired("source")
	migrateCmd.PersistentFlags().StringP("target", "t", "", "Target Kubernetes cluster")
	migrateCmd.MarkPersistentFlagRequired("target")
}
