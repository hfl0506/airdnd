import { queryOptions } from "@tanstack/react-query";
import { categoryTypeApi, ListingRecord, listingRecordsApi } from "../api/listings";
import { ListingSearchInput } from "../zod/paginate";
import { useCallback, useEffect, useState } from "react";

export const listingQueryOptions = (params: ListingSearchInput) => {
  return queryOptions({
    queryKey: ["listings"],
    queryFn: () => listingRecordsApi(params),
  });
};

export const categoryQueryOptions = () => {
  return queryOptions({
    queryKey: ["category"],
    queryFn: () => categoryTypeApi(),
  });
};

type UseInfiniteScrollProps = {
  tabId: string;
}

export const useInfiniteScroll = ({ tabId }: UseInfiniteScrollProps) => {
  const [data, setData] = useState<ListingRecord[]>([]);
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("")
  const [hasNextPage, setHasNextPage] = useState(true);
  const [page, setPage] = useState(1);
  const LIMIT = 10;

  const fetchPage = useCallback(async () => {
    setLoading(true)
    try {
      const res = await listingRecordsApi({ page, limit: LIMIT, tab_id: tabId })
      setData(prev => page === 1 ? res.data.data : [...prev, ...res.data.data]);
      const shouldFetch =
        res.data.totalPage !== undefined &&
        res.data.nextPage !== undefined &&
        res.data.nextPage <= res.data.totalPage;
      setHasNextPage(shouldFetch)
    } catch (error) {
      setErrorMessage(`${error}`)
    } finally {
      setLoading(false)
    }
  }, [page, LIMIT, tabId])

  const fetchNextPage = () => {
    setPage(prev => prev + 1)
  }

  const refetch = async () => {
    setPage(1);
    await fetchPage();
  }

  useEffect(() => {
    fetchPage();
  }, [page, tabId])

  return {
    data,
    loading,
    errorMessage,
    hasNextPage,
    fetchNextPage,
    refetch
  }
}
