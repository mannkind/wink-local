import * as React from "react";
import * as ReactDOM from "react-dom";
import { AppContainer } from "react-hot-loader";
import App from "./app/App";
import WinkStore from "./redux/Store";

import "bootstrap/dist/css/bootstrap.min.css";

const load = () => ReactDOM.render(
    <AppContainer>
        <App store={WinkStore} />
    </AppContainer>,
    document.getElementById("root"),
);

if (module.hot) {
  module.hot.accept("./app/App", load);
}

load();
