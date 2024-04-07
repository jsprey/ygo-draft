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
            <span className={labelCN}>Sets:</span>
            <span className={textCN}>{settings.sets.map((value, index) => {
                    return (index === 0 ? "" : ", ") + value.set_name
                }
            )}</span>
        </div>
    </div>
}

export default DraftSettingsDetails
