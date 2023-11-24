import { app, BrowserWindow, Menu, Tray } from 'electron';

const production = process.env.NODE_ENV === 'production';

// we'll accept these errors as they're handled early on
//@ts-expect-error
let tray: Tray = null;
//@ts-expect-error
let mainWindow: BrowserWindow = null;


app.on('ready', () => {
    mainWindow = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
            nodeIntegration: false,
        },
    });

    mainWindow.loadFile('./initial.html');
    
    tray = new Tray('../assets/Alecive-Flatwoken-Apps-Home.512.png');
    tray.setToolTip('This is my application.');
    tray.setContextMenu(
        Menu.buildFromTemplate([
            {
                label: 'Show App', click: function () {
                    mainWindow.show();
                }
            },
            {
                label: 'Quit', click: function () {
                    app.quit();
                }
            }
        ])
    );

    
});


