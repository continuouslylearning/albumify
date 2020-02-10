import DropZone from 'react-dropzone';
import React from 'react'
import { connect } from 'react-redux';

import Album from './Album';
import { getAlbum, postImages} from './actions/album';

const style = {
    height: 'auto',
    minHeight: '100%'
};

class AlbumContainer extends React.Component {
    onDrop = (files) => {
        return this.props.postImages(files)
            .then(() => this.props.getAlbum());
    }

    render = () => {
        return (
            <DropZone
                onDrop={this.onDrop}
                style={style}
            >   
                { ({getRootProps, getInputProps}) => 
                    (   
                        <div {...getInputProps({ style })} {...getRootProps()}>
                            <Album/>
                        </div>
                    )
                }   
            </DropZone>
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
        postImages: (images) => dispatch(postImages(images))
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(AlbumContainer);