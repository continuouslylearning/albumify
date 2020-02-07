import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Container from './DropComponent';
import Landing from './Landing';
import * as serviceWorker from './serviceWorker';
import { Provider } from 'react-redux';
import store from './store';

ReactDOM.render(<Provider store={store}><Landing/></Provider>, document.getElementById('root'));
serviceWorker.unregister();
