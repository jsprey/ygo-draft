import {Card,} from "../api/CardModel"
import SingleCardViewer from "../draft/shared/deck/SingleCardViewer";
import "./YgoBackground.css"
import {useRandomCards} from "../api/hooks/cards/useCards";
import {CardFilter} from "../api/CardFilter";
import React, {useEffect, useState} from "react";
import Alert from "../core/Alert";
import Spinner from "../core/Spinner";

function YgoBackground() {
    const [allowedToFetch, setAllowedToFetch] = useState<boolean>(true)
    const {data, isLoading, error} = useRandomCards("login", 180, {} as CardFilter, {
        refetchOnWindowFocus: allowedToFetch,
        enabled: allowedToFetch
    })

    useEffect(() => {
        if (data) {
            setAllowedToFetch(false)
        }
    }, [data]);

    let content
    if (isLoading) {
        content = <Spinner/>
    } else if (error) {
        content = <Alert variant={'danger'}>Failed to load background images!</Alert>
    } else if (data) {
        let myInt = 50000

        let cardsViewBody = data.cards.map((card: Card) =>
            <span key={myInt++}><SingleCardViewer className={"card"} card={card} onlyImage={true} readonly={true}/></span>
        );
        content = <div className={"blur-sm loginBackgroundContainer mySpecialBackground flex flex-wrap p-3 gap-3 bg-black -z-50 select-none"}>
            {cardsViewBody}
        </div>
    }

    return <>{content}</>
}

export default YgoBackground
