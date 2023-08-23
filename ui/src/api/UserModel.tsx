/**
 * GetUserReponse is the response from the GET /user endpoint.
 */
export type GetUserReponse = {
    id: number
    email: string
    display_name: string
    is_admin: boolean
}