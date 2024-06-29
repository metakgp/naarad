import { Component, createSignal, onMount } from "solid-js";
import { BACKEND_URL } from "../constants";
import toast from "solid-toast";
import spinner from "../assets/spinner.svg";
import check from "../assets/check.png";
import cross from "../assets/cross.png";
import { RedirectToast } from "../components/RedirectToast";

export const Register: Component = () => {
    const [getStatus, setStatus] = createSignal("Initiating User Registration");
    const [getMsg, setMsg] = createSignal("\u00A0");
    const [getIsLoad, setIsLoad] = createSignal(true);
    const [getIsErr, setIsErr] = createSignal(false);

    const PROD = import.meta.env.PROD;

    onMount(() => {
        if (!PROD) {
            setTimeout(() => {
                setMsg("This is a dummy message for development alksdfjalksdfjalskdjflaksjfdlaksjdflkasjfda");
            }, 3000);
        }
        fetch(BACKEND_URL + "/register", {
            method: "GET",
            credentials: "include",
        })
            .then((data) => {
                if (data.ok) {
                    setIsLoad(false);
                    setStatus("User Registration Successful");
                    setMsg("Credentials sent to your institute email");

                    if (PROD) {
                        toast.custom(
                            (t) => (
                                <RedirectToast
                                    duration={3000}
                                    redirect_url="https://naarad.metakgp.org/login"
                                    url_name="Naarad Login"
                                    t={t}
                                />
                            ),
                            { duration: 3000 }
                        );
                    }
                } else if (data.status === 409) {
                    setIsLoad(false);
                    setStatus("User Already Registered");
                    setMsg("Search your institute email for credentials");

                    if (PROD) {
                        toast.custom(
                            (t) => (
                                <RedirectToast
                                    duration={3000}
                                    redirect_url="https://naarad.metakgp.org/login"
                                    url_name="Naarad Login"
                                    t={t}
                                />
                            ),
                            { duration: 3000 }
                        );
                    }
                } else if (data.status == 401) {
                    setIsLoad(true);
                    setStatus("Redirecting to Heimdall");

                    if (PROD) {
                        toast.custom(
                            (t) => (
                                <RedirectToast
                                    duration={3000}
                                    redirect_url="https://heimdall.metakgp.org?redirect_url=https://naarad.metakgp.org/signup"
                                    url_name="Heimdall"
                                    t={t}
                                />
                            ),
                            { duration: 3000 }
                        );
                    }
                } else {
                    setIsLoad(false);
                    setIsErr(true);
                    setStatus("User Registration Failed");
                    data.text().then((bodyData) => {
                        setMsg(bodyData);
                    });
                }
            })
            .catch((err) => {
                setIsLoad(false);
                setIsErr(true);
                setStatus("User Registration Failed");
                setMsg(err.toString());
            });
    });

    return (
        <div class="reg">
            <div class="reg-main">
                <div class="reg-title">
                    <div class="reg-title-name">Naarad</div>
                    <div class="reg-title-desc">
                        Delivering real-time notices to KGPians
                    </div>
                    <div class="reg-status-svg">
                        {getIsLoad() == true ? (
                            <img src={spinner} />
                        ) : getIsErr() == true ? (
                            <img src={cross} />
                        ) : (
                            <img src={check} />
                        )}
                    </div>
                </div>
                <div class="reg-status">
                    <div class="reg-status-title">{getStatus()}</div>
                    <div class="reg-status-text">{`${getMsg()}`}</div>
                </div>
                <div class="reg-footer">
                    <p class="reg-footer">
                        Made with ❤️ and {"</>"}
                        <br />
                        <a
                            href="https://github.com/metakgp/naarad"
                            target="_blank"
                        >
                            metaKGP
                        </a>
                    </p>
                </div>
            </div>
        </div>
    );
};
