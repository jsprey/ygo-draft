import {UseQueryResult} from "react-query";
import {Deck} from "../../CardModel";
import {useMagicMethod} from "./useCards";

const QueryParameterCardIDs = "cards"

export function useCardsBulk(id: string[], queryOptions: any = {}): UseQueryResult<Deck> {
    let queryMap = new Map<string, string>();
    const cardIdList = id.join(",");
    queryMap.set(QueryParameterCardIDs, cardIdList)
    return useMagicMethod<Deck>(["cards", "bulk", cardIdList], `cards/bulk`, queryMap, queryOptions)
}