import { Component, createSignal } from "solid-js";
import "../styles/Instruct.scss"

export const Instruct: Component = () => {
    return (
        <div class="ins">
            <div class="ins-main">
                <div class="ins-title">
                    <div class="ins-title-name">Naarad</div>
                </div>
                <div class="ins-info">
                    <div class="ins-info-title">Desktop Setup</div>
                    <ol class="ins-info-dsktp">
                        <li>Visit <a href="https://naarad.metakgp.org/web">Naarad Portal</a> and click on signin, enter your confidentials used during registration and click on Sign In.</li>
                        <li>Click on Grant Now on the top-left side.</li>
                        <li>Click on the Subscribe to topic button present on the left side of the screen.</li>
                        <li>Enter name of the topic as naarad-cdc and click on subscribe.</li>
                    </ol>
                    <div class="ins-info-title">Mobile Setup</div>
                    <ol class="ins-info-mobile">
                        <li>Download <a href="https://play.google.com/store/apps/details?id=io.heckel.ntfy&hl=en_IN&pli=1">NTFY app</a> from play store.</li>
                        <li>Click on the plus in the bottom right corner.</li>
                        <li>Check the box corresponding to `Use another server`.</li>
                        <li>Enter name of the topic as naarad-cdc and service url as `https://naarad.metakgp.org/` and click on subscribe.</li>
                        <li>Enter the credentials used during registration and click on login.</li>
                    </ol>
                </div>
            </div>
        </div>
    )
}