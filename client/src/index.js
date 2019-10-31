import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Container from './DropComponent';
import * as serviceWorker from './serviceWorker';
import { Provider } from 'react-redux';
import store from './store';

ReactDOM.render(<Provider store={store}><Container/></Provider>, document.getElementById('root'));
serviceWorker.unregister();
