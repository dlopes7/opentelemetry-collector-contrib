// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	conventionsv112 "go.opentelemetry.io/collector/semconv/v1.12.0"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func makeService(resource pcommon.Resource) *awsxray.ServiceData {
	var service *awsxray.ServiceData

	verStr, ok := resource.Attributes().Get(conventionsv112.AttributeServiceVersion)
	if !ok {
		verStr, ok = resource.Attributes().Get(conventionsv112.AttributeContainerImageTag)
	}
	if ok {
		service = &awsxray.ServiceData{
			Version: awsxray.String(verStr.Str()),
		}
	}
	return service
}
