import { Routes } from '@angular/router';
import { NotFound } from './pages/not-found/not-found';

/*
to create a route, you need to add in Routes the following json object :

{
    path: 'url',
    component: component,
}

you also need to import the required component
*/

export const routes: Routes = [
    {
        path: '**', // all other paths
        component: NotFound,
    }
];
