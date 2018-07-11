all: apply

init: ## Init terraform for usage.
	terraform init -input=false

plan: init ## Plan infrastructure/terraform operations.
	terraform plan

apply: init plan ## Apply infrastructure plan.
	terraform apply

destroy: init ## Destroy infrastructure.
	terraform destroy

help: ## Show this help menu.
	@grep -E '^[a-zA-Z_-%]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all init plan apply destroy help
