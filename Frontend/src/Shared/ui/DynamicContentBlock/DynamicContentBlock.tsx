import { useEffect } from 'react';
import {  TypeOur } from '../../types/TypeDinamicBlock';
import './DynamicContentBlock.scss'


const DynamicContentBlock = ({ classes, data}: TypeOur) => {

    return (
        <div className='DynamicContentBlock'>
          <h2 className={'DynamicContentBlock__title ' + classes?.classtitle}>{data.title}</h2>
          <p className={'DynamicContentBlock__subtitle ' + classes?.subtitle}>{data.description}</p>
          <button className={'DynamicContentBlock__buttonText ' + classes?.textbutton}>{data.buttonText}</button>
        </div>
      );
}

export {DynamicContentBlock}

