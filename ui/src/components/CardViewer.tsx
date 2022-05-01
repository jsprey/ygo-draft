import {Card, useCard} from "./hooks/useRandomCard";

export type CardViewerProps = {
    id: string
}

function CardViewer(props: CardViewerProps) {
    const {isLoading, error, data} = useCard(props.id)

    let body;
    if (isLoading) {
        body = <p>Loading...</p>
    } else if (!data) {
        body = <p>Fehler beim Laden der Karte!</p>
    } else if (data && data.card_images) {
        console.log("My api: " + JSON.stringify(data))
        body = getCardAsImage(data)
    } else {
        body = <p>no data</p>
    }

    return <div className={"w-36"}>
        {body}
    </div>
}

function getCardAsImage(data: Card): JSX.Element {
    return <img
        src={data.card_images[0].image_url_small}
        alt="new"
    />
}

function getCard(data: Card): JSX.Element {
    return <div className="md:col-span-2 lg:col-span-1">
        <div className="h-full py-8 px-6 space-y-6 rounded-xl border border-gray-200 bg-white">
            <div>
                <p className="text-l text-gray-600 text-center">{data.id}</p>
                <p className="text-xl text-gray-600 text-center">{data.name}</p>
            </div>
            <table className="w-full text-gray-600">
                <tbody>
                <tr>
                    <td className="py-2">Type</td>
                    <td className="text-gray-500">{data.type}</td>
                </tr>
                <tr>
                    <td className="py-2">ATK</td>
                    <td className="text-gray-500">{data.atk}</td>
                </tr>
                <tr>
                    <td className="py-2">DEF</td>
                    <td className="text-gray-500">{data.def}</td>
                </tr>
                </tbody>
            </table>

            <p className="text-m text-gray-400 text-center">{data.desc}</p>
        </div>
    </div>
}

export default CardViewer