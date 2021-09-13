import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'Development streamlined',
    description: (
      <>
        Ox aims to facilitate development tasks. It is a tool that helps you improve your team release cycle by automating repetitive tasks.
      </>
    ),
  },
  {
    title: 'Maintainability in mind',
    description: (
      <>
        Besides allowing to accelerate your development, Ox aims to make your code more maintainable. It is thought not only for MVP's but large codebases.
      </>
    ),
  },
  {
    title: 'Plugin System',
    description: (
      <>
        Extend or customize your development environment with plugins. Ox allows to plug the tools you need to get the job done.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        {/* <Svg className={styles.featureSvg} alt={title} /> */}
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
