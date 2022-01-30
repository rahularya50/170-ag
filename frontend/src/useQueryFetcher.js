// @flow

import { useState } from "react";
import { fetchQuery, useRelayEnvironment } from "react-relay";
import type {
  GraphQLTaggedNode,
  OperationType,
  VariablesOf,
} from "relay-runtime";

export function useQueryFetcher<T: OperationType>(
  operation: GraphQLTaggedNode,
  params: VariablesOf<T>,
  onComplete: (T["response"]) => mixed,
  onError: (Error) => mixed
): [() => void, boolean] {
  const environment = useRelayEnvironment();
  const [isLoading, setIsLoading] = useState(false);
  const fetch = () => {
    setIsLoading(true);
    fetchQuery<T>(environment, operation, params).subscribe({
      next: (e) => {
        setIsLoading(false);
        onComplete(e);
      },
      error: onError,
    });
  };
  return [fetch, isLoading];
}
