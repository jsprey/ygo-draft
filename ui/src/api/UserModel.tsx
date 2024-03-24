/**
 * GetUserReponse is the response from the GET /user endpoint.
 */
export type GetUserReponse = {
    id: number
    email: string
    display_name: string
    is_admin: boolean
}

export type Friend = {
    id: number
    name: string
}

export type FriendRequest = {
    id: number
    name: string
    invitation_date: Date
}
/**
 * GetFriendsResponse is the response from the GET /user/friends endpoint.
 */
export type GetFriendsResponse = Friend[]

/**
 * GetFriendRequestsResponse is the response from the GET /user/friends/requests endpoint.
 */
export type GetFriendRequestsResponse = FriendRequest[]

export function getTimeDifferenceString(date: Date): string {
    const currentDate = Date.now()
    const timeDifference = currentDate - new Date(date).getTime();
    const secondsDifference = Math.floor(timeDifference / 1000);
    const minutesDifference = Math.floor(secondsDifference / 60);
    const hoursDifference = Math.floor(minutesDifference / 60);
    const daysDifference = Math.floor(hoursDifference / 24);
    const weeksDifference = Math.floor(daysDifference / 7);
    const monthsDifference = Math.floor(daysDifference / 30);
    const yearsDifference = Math.floor(daysDifference / 365);

    if (yearsDifference > 0) {
        return `${yearsDifference} year${yearsDifference > 1 ? 's' : ''} ago`;
    } else if (monthsDifference > 0) {
        return `${monthsDifference} month${monthsDifference > 1 ? 's' : ''} ago`;
    } else if (weeksDifference > 0) {
        return `${weeksDifference} week${weeksDifference > 1 ? 's' : ''} ago`;
    } else if (daysDifference > 0) {
        return `${daysDifference} day${daysDifference > 1 ? 's' : ''} ago`;
    } else if (hoursDifference > 0) {
        return `${hoursDifference} hour${hoursDifference > 1 ? 's' : ''} ago`;
    } else if (minutesDifference > 0) {
        return `${minutesDifference} minute${minutesDifference > 1 ? 's' : ''} ago`;
    } else {
        return `${secondsDifference} second${secondsDifference !== 1 ? 's' : ''} ago`;
    }
}