export type SearchParams = {
  type: string;
};

export type GenericRes<T> = {
  data: T;
  success: boolean;
  total?: number;
  nextPage?: number;
  totalPage?: number;
};
