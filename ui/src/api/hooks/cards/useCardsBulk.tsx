import {useQuery, UseQueryResult} from "react-query";
import {Card, Deck} from "../../CardModel";
import {PUBLIC_URL} from "../../../index";
import {QueryKey} from "react-query/types/core/types";
import {CardFilter, FilterToQuery} from "../../CardFilter";
import axios from "axios";
import {useMagicMethod} from "./useCards";

const QueryParameterCardIDs = "cards"

export function useCardsBulk(id: string[], queryOptions: any = {}): UseQueryResult<Deck> {
    let queryMap = new Map<string, string>();
    const cardIdList = id.join(",");
    queryMap.set(QueryParameterCardIDs, cardIdList)
    return useMagicMethod<Deck>(["cards", "bulk", cardIdList], `cards/bulk`, queryMap, queryOptions)
}