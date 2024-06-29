import { Component, createEffect, createSignal, onCleanup } from "solid-js";
import { Toast } from "solid-toast";

type IProps = {
    t: Toast;
    duration: number;
    redirect_url: string;
    url_name?: string;
};

export const RedirectToast: Component<IProps> = (props) => {
    const [life, setLife] = createSignal(100);
    const startTime = Date.now();

    createEffect(() => {
        if (props.t.paused) return;
        const interval = setInterval(() => {
            const diff = Date.now() - startTime - props.t.pauseDuration;
            setLife((diff / props.duration) * 100);
        });
        
        onCleanup(() => {
            window.location.href = props.redirect_url;
            clearInterval(interval);
        });
    });

    return (
        <div class="toast-notification toast-visible">
            <div class="toast-content">
                <div class="toast-text-content">
                    <div class="toast-title">
                        Redirecting to{" "}
                        <a href={props.redirect_url}>
                            {props.url_name
                                ? props.url_name
                                : props.redirect_url}
                        </a>
                    </div>
                </div>
            </div>
            <div class="toast-progress-bar">
                <div class="toast-progress-background"></div>
                <div
                    class="toast-progress-foreground"
                    style={{ width: `${life()}%` }}
                ></div>
            </div>
        </div>
    );
};
