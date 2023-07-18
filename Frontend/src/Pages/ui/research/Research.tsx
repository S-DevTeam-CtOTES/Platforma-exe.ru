import React from "react";
import { Exploration, Header } from "@/Widgets";
import "./research.scss";
import { App } from "@Shared/consts/index";
import Image from "@Shared/assets/img/research.svg";
import Warning from "@Shared/assets/img/Warning.svg";
import Telegram from "@Shared/assets/img/Telegram.svg";
import VK from "@Shared/assets/img/VK.svg";
import Youtube from "@Shared/assets/img/Youtube.svg";

const Research = () => {
  return (
    <section className="Research__component">
      <Header />

      <div className="container">
        <div className="Research__component-leftPart">
          <div className="warning"><img src={Warning}></img></div>
          <h1 className="Research__component-leftPart-title">В разработке</h1>
          <p className="Research__component-leftPart-subtitle">Эта страница находиться в разработке!<br></br>Чтобы быть в курсе всех последний новостей,<br></br> приглашаем вас следить за нами в социальных сетях</p>
          <div className="Research__component-leftPart-links">
            <a href="https://t.me/platformaexe"><img src={Telegram}></img></a>
            <a href="https://vk.com/platformaexe"><img src={VK}></img></a>
            <a href="https://www.youtube.com/@platforma-exe"><img src={Youtube}></img></a>
          </div>
        </div>
          
        <div className="Research__component-pic">
          <img src={Image}></img>
        </div>

      </div>
      <div className="Research__component-cop">{App.Copyright}</div>
    </section>
  );
};

export { Research };

{
  /* <div className="Error__404-img">
                <img src={Oops404}></img>
        </div>  

        <div className="Error__404-cop">
            {App.Copyright}
  </div> */
}
