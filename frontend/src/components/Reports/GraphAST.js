import React,{useEffect} from 'react';
import Graphviz from 'graphviz-react';
function Graph({dot}){
    /*useEffect(() => {
        d3.select("#graph-body").graphviz().renderDot(dot);
        zoom:options;
    }, [dot]);
    return ( <div
        id="graph-body"
        style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
        }}
    />);*/
    return <Graphviz dot={dot} options={{
        zoom: true,
        width: window.innerWidth,
        height: window.innerHeight,
        engine: "dot"
      }} />;
};

export default Graph;