all: apply

init: ## Init terraform for usage.
	terraform init -input=false

plan: init ## Plan infrastructure/terraform operations.
	terraform plan -var-file=vars.tfvars

apply: init plan ## Apply infrastructure plan.
	terraform apply -var-file=vars.tfvars

destroy: init ## Destroy infrastructure.
	terraform destroy -var-file=vars.tfvars

help: ## Show this help menu.
	@grep -E '^[a-zA-Z_-%]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all init plan apply destroy help
