import React from 'react';
import computer from '@Shared/assets/img/computer.svg'

import {DynamicContentBlock} from '@/Shared';

import './Exploration.scss'

const Exploration = () => {


  const data = {
    title: "ИССЛЕДОВАНИЕ",
    description: "В современном обществе существует проблема профориентации, команда проекта пытается её решить! Мы подготовили профориентационный тест, который поможет определиться с наиболее подходящей для вас профессией и помочь в её освоении.",
    buttonText: "ПОДРОБНЕЕ",
    
  };

  const classes = {
    classtitle: 'Exploration__title',
    subtitle: 'Exploration__subtitle',
    textbutton: 'Exploration__textbutton'
  }

  return (
    <section className='Exploration'>
        <div className="container">
            <div className="Exploration__wrapper">
            <div className="Exploration__img"><img src={computer} alt="computer" /></div>
              <DynamicContentBlock classes={classes} data={data}/>
            </div>
        </div>
    </section>
  )
}

export {Exploration}