import React, {ReactNode} from "react";
import {DraftSettings} from "../../../api/Draft";
import classNames from "classnames";

type DraftSettingsDetailsProps = {
    settings: DraftSettings
    containerClasses?: string
    textClasses?: string
    labelClasses?: string
};

function DraftSettingsDetails(props: DraftSettingsDetailsProps) {
    const settings = props.settings

    const containerCN = classNames(props.containerClasses)
    const gridCN = classNames("grid grid-cols-8")
    const labelCN = classNames("col-span-2", "fw-bold", props.labelClasses)
    const textCN = classNames("col-span-6", props.textClasses)

    function getSets(): ReactNode {
        return <>
            <div className={"p-2 mb-0 mt-2 font-bold bg-light-3 dark:bg-dark-3"}>
                Sets
            </div>
            <div style={{maxHeight: "15rem"}}
                 className={classNames(textCN, "overflow-y-auto p-2 mb-1 bg-light-1 dark:bg-dark-1")}>
                {props.settings.sets.map(value => {
                    return <li key={`draft-settings-details-set-${value.set_code}`}>{value.set_name}</li>
                })}
            </div>
        </>
    }

    return <div className={containerCN}>
        <div className={gridCN}>
            <span className={labelCN}>Mode:</span>
            <span
                className={textCN}>{settings.mode === "bestof" ? `Best of ${settings.mode_value} Rounds` : `${settings.mode_value} Rounds`}</span>
            <span className={labelCN}>Main Deck Drafts:</span>
            <span className={textCN}>{settings.main_deck_draws}</span>
            <span className={labelCN}>Main Deck Drafts Size:</span>
            <span className={textCN}>{settings.main_deck_size}</span>
            <span className={labelCN}>Extra Deck Drafts:</span>
            <span className={textCN}>{settings.extra_deck_draws}</span>
            <span className={labelCN}>Extra Deck Drafts Size:</span>
            <span className={textCN}>{settings.extra_deck_size}</span>
        </div>
        {props.settings.sets.length === 0 ? <span className={"font-bold"}>All cards allowed!</span> : getSets()}
    </div>
}

export default DraftSettingsDetails
