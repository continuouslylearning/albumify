import DropZone from 'react-dropzone';
import React from 'react'
import { connect } from 'react-redux';

import Album from './Album';
import { getAlbum, postImages} from './actions/album';


class AlbumContainer extends React.Component {
    onDrop = (files) => {
        return this.props.postImages(files)
            .then(() => this.props.getAlbum());
    }

    render = () => {
        return (
            <div>
                <DropZone
                    onDrop={this.onDrop}
                >   
                    { ({getRootProps, getInputProps}) => 
                        (   
                            <div {...getRootProps()}>
                                {/* <input {...getInputProps()}/> */}
                                <Album/>
                            </div>
                        )
                    }   
                </DropZone>
            </div>
        );
    }
}

const mapStateToProps = (state) => {
    return {

    };
};

const mapDispatchToProps = (dispatch) => {
    return {
        getAlbum: () => dispatch(getAlbum()),
        postImages: (images) => dispatch(postImages(images)),
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(AlbumContainer);