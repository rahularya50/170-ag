// @flow

import { useCallback, useEffect, useState } from "react";

export function useDebounced<T>(value: T, delay: number): T {
  const [debouncedValue, setDebouncedValue] = useState(value);
  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value);
    }, delay);
    return () => clearTimeout(handler);
  }, [value, delay]);
  return debouncedValue;
}

/*
  Will invoke callback `delay` ms after its last invocation.
*/
export function useDebouncedCallback<T: (...Array<any>) => void>(
  callback: T,
  delay: number
): T {
  const [pendingClosure, setPendingClosure] = useState();

  useEffect(() => {
    const handler = setTimeout(() => {
      if (pendingClosure) pendingClosure();
    }, delay);
    return () => clearTimeout(handler);
  }, [pendingClosure, delay]);

  // $FlowFixMe: Needed to make callers type-safe
  return useCallback(
    (...args) => {
      setPendingClosure(() => () => {
        callback(...args);
      });
    },
    [callback]
  );
}
