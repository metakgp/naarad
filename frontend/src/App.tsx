import type { Component } from "solid-js";
import { Router, Route } from "@solidjs/router";
import { Toaster } from "solid-toast";
import { Register } from "./pages/Register";
import { Guide } from "./pages/Guide";
import "./styles/index.scss";

const App: Component = () => {
    return (
        <>
            <Toaster position="bottom-center" gutter={8} />
            <Router>
                <Route path="/guide" component={Guide} />
                <Route path="*404" component={Register} />
            </Router>
        </>
    );
};

export default App;
