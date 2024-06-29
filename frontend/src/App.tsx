import type { Component } from "solid-js";
import { Register } from "./pages/Register";
import "./styles/index.scss";
import { Toaster } from "solid-toast";

const App: Component = () => {
    return (
        <>
            <Toaster position="bottom-center" gutter={8} />
            <Register />
        </>
    );
};

export default App;
