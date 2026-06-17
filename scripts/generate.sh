#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GEN_IMAGE="openapitools/openapi-generator-cli:v7.13.0"
CONFIG_FILE="/local/scripts/openapi-generator-config.yaml"

run_generate() {
  local spec_path="$1"
  local out_dir="$2"
  local pkg_name="$3"

  rm -rf "${ROOT_DIR}/${out_dir}"
  docker run --rm \
    -u "$(id -u):$(id -g)" \
    -v "${ROOT_DIR}:/local" \
    "${GEN_IMAGE}" generate \
      -c "${CONFIG_FILE}" \
      -i "/local/${spec_path}" \
      -o "/local/${out_dir}" \
      --additional-properties "packageName=${pkg_name}"

  rm -rf "${ROOT_DIR}/${out_dir}/docs" \
         "${ROOT_DIR}/${out_dir}/.openapi-generator"
  rm -f "${ROOT_DIR}/${out_dir}/.openapi-generator-ignore" \
        "${ROOT_DIR}/${out_dir}/README.md" \
        "${ROOT_DIR}/${out_dir}/git_push.sh" \
        "${ROOT_DIR}/${out_dir}/.travis.yml" \
        "${ROOT_DIR}/${out_dir}/go.mod" \
        "${ROOT_DIR}/${out_dir}/go.sum"
}

run_generate "specs/wb/01-general.yaml" "internal/generated/general" "wbgeneral"
run_generate "specs/wb/02-products.yaml" "internal/generated/products" "wbproducts"
run_generate "specs/wb/06-reports.yaml" "internal/generated/reports" "wbreports"
run_generate "specs/wb/07-analytics.yaml" "internal/generated/analytics" "wbanalytics"
run_generate "specs/wb/08-orders-fbw.yaml" "internal/generated/ordersfbw" "wbordersfbw"
run_generate "specs/wb/09-in-store-pickup.yaml" "internal/generated/clickcollect" "wbclickcollect"
run_generate "specs/wb/04-orders-dbw.yaml" "internal/generated/dbw" "wbdbw"
run_generate "specs/wb/05-orders-dbs.yaml" "internal/generated/dbs" "wbdbs"

# WB FBS spec currently contains a schema name starting with digits.
# OpenAPI generator emits invalid Go identifiers for the derived inline models,
# so we normalize that component name and its refs in a temporary working copy.
FBS_WORK_SPEC="specs/wb/.03-orders-fbs.work.yaml"
cp "${ROOT_DIR}/specs/wb/03-orders-fbs.yaml" "${ROOT_DIR}/${FBS_WORK_SPEC}"
perl -pi -e 's#/components/schemas/409SupplyDeliverError#/components/schemas/Model409SupplyDeliverError#g' "${ROOT_DIR}/${FBS_WORK_SPEC}"
perl -pi -e 's/^\s{4}409SupplyDeliverError:/    Model409SupplyDeliverError:/g' "${ROOT_DIR}/${FBS_WORK_SPEC}"
run_generate "${FBS_WORK_SPEC}" "internal/generated/fbs" "wbfbs"
rm -f "${ROOT_DIR}/${FBS_WORK_SPEC}"
