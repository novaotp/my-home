export interface APIResponse {
    success: boolean,
    message: string
}

export type WithData<T, D> = T & { data: D };
