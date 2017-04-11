import * as React from "react";
import * as ReactDOM from "react-dom";
import App from "./app/App";
import WinkStore from "./redux/Store";

import "bootstrap/dist/css/bootstrap.min.css";

const load = () => {
    ReactDOM.render(
        <App store={WinkStore} />,
        document.getElementById("root"),
    );
};

load();
