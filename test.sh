#!/bin/sh

docker run --rm \
   -v $(pwd):$(pwd) \
   -w $(pwd) \
   jmccann/drone-terraform --actions validate --actions plan \
   --tf.version 0.12.8 \
   --root_dir terraform \
   --init_options '{"backend-config": ["config/dev/backend.tfvars"] }' \
   --var_files config/dev/terraform.tfvars \
   --env-file .env \
   --role_arn_to_assume arn:aws:iam::374047294805:role/automation-drone
