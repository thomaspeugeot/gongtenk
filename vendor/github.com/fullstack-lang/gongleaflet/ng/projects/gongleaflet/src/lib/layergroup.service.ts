// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { LayerGroupDB } from './layergroup-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class LayerGroupService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LayerGroupServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private layergroupsUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.layergroupsUrl = origin + '/api/github.com/fullstack-lang/gongleaflet/go/v1/layergroups';
  }

  /** GET layergroups from the server */
  getLayerGroups(): Observable<LayerGroupDB[]> {
    return this.http.get<LayerGroupDB[]>(this.layergroupsUrl)
      .pipe(
        tap(_ => this.log('fetched layergroups')),
        catchError(this.handleError<LayerGroupDB[]>('getLayerGroups', []))
      );
  }

  /** GET layergroup by id. Will 404 if id not found */
  getLayerGroup(id: number): Observable<LayerGroupDB> {
    const url = `${this.layergroupsUrl}/${id}`;
    return this.http.get<LayerGroupDB>(url).pipe(
      tap(_ => this.log(`fetched layergroup id=${id}`)),
      catchError(this.handleError<LayerGroupDB>(`getLayerGroup id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new layergroup to the server */
  postLayerGroup(layergroupdb: LayerGroupDB): Observable<LayerGroupDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.post<LayerGroupDB>(this.layergroupsUrl, layergroupdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted layergroupdb id=${layergroupdb.ID}`)
      }),
      catchError(this.handleError<LayerGroupDB>('postLayerGroup'))
    );
  }

  /** DELETE: delete the layergroupdb from the server */
  deleteLayerGroup(layergroupdb: LayerGroupDB | number): Observable<LayerGroupDB> {
    const id = typeof layergroupdb === 'number' ? layergroupdb : layergroupdb.ID;
    const url = `${this.layergroupsUrl}/${id}`;

    return this.http.delete<LayerGroupDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted layergroupdb id=${id}`)),
      catchError(this.handleError<LayerGroupDB>('deleteLayerGroup'))
    );
  }

  /** PUT: update the layergroupdb on the server */
  updateLayerGroup(layergroupdb: LayerGroupDB): Observable<LayerGroupDB> {
    const id = typeof layergroupdb === 'number' ? layergroupdb : layergroupdb.ID;
    const url = `${this.layergroupsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.put<LayerGroupDB>(url, layergroupdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated layergroupdb id=${layergroupdb.ID}`)
      }),
      catchError(this.handleError<LayerGroupDB>('updateLayerGroup'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
