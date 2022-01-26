// @flow

import type {
  VariablesOf,
  OperationType,
  GraphQLTaggedNode,
} from "relay-runtime";

import { useCallback, useState } from "react";
import { useRelayEnvironment, fetchQuery } from "react-relay";
import useInterval from "use-interval";

export function useRefreshingQuery<T: OperationType>(
  operation: GraphQLTaggedNode,
  params: VariablesOf<T>
) {
  const environment = useRelayEnvironment();
  const [isRefreshing, setIsRefreshing] = useState(false);
  const refresh = useCallback(() => {
    if (isRefreshing) {
      return;
    }
    setIsRefreshing(true);
    fetchQuery<T>(environment, operation, params).subscribe({
      complete: () => {
        setIsRefreshing(false);
      },
      error: () => setIsRefreshing(false),
    });
  }, [environment, isRefreshing, params, operation]);

  useInterval(refresh, 5000);
}
