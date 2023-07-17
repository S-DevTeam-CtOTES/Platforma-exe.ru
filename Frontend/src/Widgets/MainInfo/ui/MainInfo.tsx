import React from 'react'
import phone from '@Shared/assets/img/Phone.svg'

import { DynamicContentBlock } from '@/Shared';
import './MainInfo.scss';

const MainInfo = () => {

  const data = {
    title: "БОЛЬШЕ ЧЕМ ПЛАТФОРМА",
    description: "Образовательный онлайн-ресурс, с тестированием на профориентацию. Мы поможем определиться с профессией, которая подходит именно вам.",
    buttonText: "ПРОЙТИ ТЕСТ"
  }

  const classes = {
    classtitle: 'MainInfo__title',
    subtitle: 'MainInfo__subtitle',
    textbutton: 'MainInfo__textbutton'
  }


  return (
    <section className='MainInfo'>
        <div className="container">
          <div className="MainInfo__wrapper">
              <DynamicContentBlock classes={classes} data={data} />
              <div className="MainInfo__img"><img src={phone} alt="test" /></div>
          </div>
        </div>
    </section>
  )
}

export {MainInfo}