import * as React from "react";
import {ThumbVideo} from "../media/thumb-video";
import "./home-page.scss"
import {IMediaResponse} from "./media/media-list-page";

interface IProps {
    media: IMediaResponse[];
}

export class HomePage extends React.Component<IProps, {}> {

    public render(): JSX.Element {
        return <div id={"home-page"}>
            <div className={"row"}>
                {this.props.media.map((media: IMediaResponse): JSX.Element => {
                    return <div key={media.id} className={"col col-sm-3"}>
                        <ThumbVideo isLink={true} showMetadata={true} media={media}/>
                    </div>;
                })}
            </div>
        </div>;
    }

}
