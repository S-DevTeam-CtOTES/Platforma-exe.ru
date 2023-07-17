import { Exploration, Header } from "@/Widgets";
import React from "react";
import Oops404 from "@Shared/assets/img/404.svg";
import {App} from "@Shared/consts/index"
import "./error.scss"

const Error404 = () => {
  return (
    <section className="Error__404">
      <Header/>
      <div className="container">
        <div className="Error__404-img">
                <img src={Oops404}></img>
        </div>  

        <div className="Error__404-cop">
            {App.Copyright}
        </div>
      </div>
    </section>
  );
};

export { Error404 };
