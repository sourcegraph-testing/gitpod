/**
 * Copyright (c) 2021 Gitpod GmbH. All rights reserved.
 * Licensed under the MIT License. See License-MIT.txt in the project root for license information.
 */

{
  prometheusAlerts+:: {
    groups+: [
      {
        name: 'gitpod-component-webapp-usage-alerts',
        rules: [
          {
            alert: 'GitpodUsageScheduledReconciliationFailures',
            expr: 'sum(increase(gitpod_usage_reconcile_completed_duration_seconds_count{outcome!="success"}[1h])) > 1',
            'for': '30m',
            labels: {
              severity: 'warning',
              team: 'webapp'
            },
            annotations: {
              runbook_url: 'https://github.com/gitpod-io/runbooks/blob/main/runbooks/GitpodUsageScheduledReconciliationFailures.md',
              summary: 'There are failed scheduled reconciliations in the usage component.',
              description: 'We have accumulated {{ printf "%.2f" $value }} failures. This affects how stale usage data is and/or updating invoices in Stripe.',
            },
          },
        ],
      },
    ],
  },
}
