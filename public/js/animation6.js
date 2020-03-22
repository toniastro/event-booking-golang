/* -----------------------------------
   Code by Brandio
------------------------------------*/
"use strict"
window.onload = function() {
    var canvas = document.getElementById('pagebg');
    paper.setup(canvas);

    var rect = new paper.Path.Rectangle(0, 0, paper.view.bounds.width, paper.view.bounds.height);
    rect.fillColor = {
        gradient: {
            stops: ['#42287C', '#714091']
        },
        origin: paper.view.bounds.topLeft,
        destination: paper.view.bounds.bottomLeft
    }

    var svg = [paper.project.importSVG('<path d="M32.842,244.769c5.15,52.94,33.761,102.955,76.392,133.374,42.345,30.711,98.137,41.533,148.779,28.664,47.781-11.992,90.126-43.581,139.337-47.968,42.059-4.095,83.259,12.284,125.318,14.917,80.97,4.972,162.513-46.505,194.271-122.552,5.722-14.039,9.728-31.3,1.144-43.58-5.722-8.19-15.736-12.284-25.178-15.5C522.382,130.407,351.572,68.4,175.612,27.452c-12.875-2.925-27.181-5.85-39.2.585-11.158,5.85-17.167,18.427-22.317,30.126C83.77,126.312,53.728,194.169,23.4,262.026"/>'),
    paper.project.importSVG('<path d="M134.184,753.713c60.942,32.174,133.9,39.778,199.993,20.767S458.636,710.425,494.4,650.466c44.348-74.292,54.362-168.765,109.582-235.159,50.356-60.837,129.9-87.453,206.574-102.663s156.218-21.936,227.174-54.987,133.9-100.908,133.043-180.757L94.7,106.441Z" />'),
    paper.project.importSVG('<path d="M242.321,574.734c-43.2,1.462-78.681-33.636-89.554-52.062s-16.595-41.533-9.728-62.007c12.589-38.608,59.8-51.478,99.854-51.185,39.77.293,82.973,6.727,116.734-15.209,20.886-13.454,34.906-36.561,56.078-50.015,41.773-26.616,103-4.387,125.89,40.363S552.754,487.28,520.137,525.3c-32.331,38.023-81.542,57.912-130.754,63.177-49.212,4.972-98.709-3.217-147.062-13.747"/>'),
    paper.project.importSVG('<path d="M381.5,667.415c9.728-52.648,52.645-93.6,100.426-115.532s100.712-28.371,152.5-37.438S739.142,493.094,781.487,461.8c36.623-26.909,62.659-65.809,90.412-102.078,28.039-35.976,60.656-71.367,103.287-85.991s97.565-1.755,117.879,39.486l-16.881,428.2Z"/>')];

    var path = [new paper.Path(),new paper.Path(),new paper.Path(),new paper.Path()];

    var orgSeg = [[],[],[],[]];

    var pathOptions = [{color: '#2D1E56',opacity: 0.35,angle: 0, position: 'TL',x: -250,y: -220,scale: 1},
                   {color: '#2D1E56',opacity: 0.18,angle: 0, position: 'TL',x: -50,y: -50,scale: 1},
                   {color: '#2D1E56',opacity: 0.18,angle: 0, position: 'BL',x: 100,y: 130,scale: 1},
                   {color: '#2D1E56',opacity: 0.18,angle: 0, position: 'BR',x: 150,y: 160,scale: 1}];

    var moveSize = 10;
    var moveSpeed = 300;

    initializePath();

    function initializePath() {
        for(var p=0;p<path.length;p++){
            path[p].segments = [];
            path[p] = svg[p];

            path[p].fillColor = pathOptions[p].color;
            path[p].fillColor.alpha = pathOptions[p].opacity;

            path[p].scale(pathOptions[p].scale);
            path[p].rotate(pathOptions[p].angle);
        
            fixSize();
            
            //path[p].smooth({ type: 'continuous' });

            for (var i = 0; i < path[p].segments.length; i++) {
                orgSeg[p].push([path[p].segments[i].point.x,path[p].segments[i].point.y]);
            }
        }
    }
    function fixSize(){
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;
        paper.view.bounds.width = canvas.width;
        paper.view.bounds.height = canvas.height;
        rect.bounds.width = paper.view.bounds.width;
        rect.bounds.height = paper.view.bounds.height;

        for(var p=0;p<path.length;p++){
            var posX = 0;
            var posY = 0;
            if(pathOptions[p].position=="TL"){
                posX = paper.view.bounds.topLeft.x + path[p].bounds.width/2;
                posY = paper.view.bounds.topLeft.y + path[p].bounds.height/2;
            }else if(pathOptions[p].position=="BL"){
                posX = paper.view.bounds.bottomLeft.x + path[p].bounds.width/2;
                posY = paper.view.bounds.bottomLeft.y - path[p].bounds.height/2;
            }else if(pathOptions[p].position=="TR"){
                posX = paper.view.bounds.topRight.x - path[p].bounds.width/2;
                posY = paper.view.bounds.topRight.y + path[p].bounds.height/2;
            }else if(pathOptions[p].position=="BR"){
                posX = paper.view.bounds.bottomRight.x - path[p].bounds.width/2;
                posY = paper.view.bounds.bottomRight.y - path[p].bounds.height/2;
            }
            path[p].position = new paper.Point(posX + pathOptions[p].x,posY + pathOptions[p].y);
        }
    }
    paper.view.onFrame = function(event) {
        for(var p=0;p<path.length;p++){
            for (var i = 0; i < path[p].segments.length; i++) {
                var xPos = Math.cos(event.count*(i+1)/moveSpeed)*moveSize;
                var yPos = Math.sin(event.count*(i+1)/moveSpeed)*moveSize;
                path[p].segments[i].point.x = orgSeg[p][i][0] + xPos;
                path[p].segments[i].point.y = orgSeg[p][i][1] + yPos;
            }
        }
    }
    paper.view.onResize = function(event) {
        fixSize();
    }
}