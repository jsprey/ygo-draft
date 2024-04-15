import React from "react";
import {DraftSettings} from "../../api/Draft";
import classNames from "classnames";

type DraftSettingsDetailsProps = {
    settings: DraftSettings
    containerClasses?: string
    textClasses?: string
    labelClasses?: string
};

function DraftSettingsDetails(props: DraftSettingsDetailsProps) {
    const settings = props.settings

    const containerCN = classNames("grid grid-cols-8", props.containerClasses)
    const labelCN = classNames("col-span-2", "fw-bold text-neutral-900 dark:text-neutral-50", props.labelClasses)
    const textCN = classNames("col-span-6", "text-neutral-900 dark:text-neutral-50", props.textClasses)

    return <div>
        <div className={containerCN}>
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
        <div className={"p-2 mb-0 mt-2 font-bold bg-gray-400 dark:bg-gray-700 dark:text-white"}>
            Sets
        </div>
        <div style={{maxHeight: "15rem"}}
             className={classNames(textCN, "overflow-y-auto p-2 mb-1 bg-gray-200 dark:bg-gray-600 dark:text-white")}>
            {props.settings.sets.map(value => {
                return <li key={`draft-settings-details-set-${value.set_code}`}>{value.set_name}</li>
            })}
        </div>
    </div>
}

export default DraftSettingsDetails
