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

import { IndividualDB } from './individual-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class IndividualService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  IndividualServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private individualsUrl: string

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
    this.individualsUrl = origin + '/api/github.com/thomaspeugeot/gongtenk/go/v1/individuals';
  }

  /** GET individuals from the server */
  getIndividuals(): Observable<IndividualDB[]> {
    return this.http.get<IndividualDB[]>(this.individualsUrl)
      .pipe(
        tap(_ => this.log('fetched individuals')),
        catchError(this.handleError<IndividualDB[]>('getIndividuals', []))
      );
  }

  /** GET individual by id. Will 404 if id not found */
  getIndividual(id: number): Observable<IndividualDB> {
    const url = `${this.individualsUrl}/${id}`;
    return this.http.get<IndividualDB>(url).pipe(
      tap(_ => this.log(`fetched individual id=${id}`)),
      catchError(this.handleError<IndividualDB>(`getIndividual id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new individual to the server */
  postIndividual(individualdb: IndividualDB): Observable<IndividualDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.post<IndividualDB>(this.individualsUrl, individualdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted individualdb id=${individualdb.ID}`)
      }),
      catchError(this.handleError<IndividualDB>('postIndividual'))
    );
  }

  /** DELETE: delete the individualdb from the server */
  deleteIndividual(individualdb: IndividualDB | number): Observable<IndividualDB> {
    const id = typeof individualdb === 'number' ? individualdb : individualdb.ID;
    const url = `${this.individualsUrl}/${id}`;

    return this.http.delete<IndividualDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted individualdb id=${id}`)),
      catchError(this.handleError<IndividualDB>('deleteIndividual'))
    );
  }

  /** PUT: update the individualdb on the server */
  updateIndividual(individualdb: IndividualDB): Observable<IndividualDB> {
    const id = typeof individualdb === 'number' ? individualdb : individualdb.ID;
    const url = `${this.individualsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.put<IndividualDB>(url, individualdb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated individualdb id=${individualdb.ID}`)
      }),
      catchError(this.handleError<IndividualDB>('updateIndividual'))
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
