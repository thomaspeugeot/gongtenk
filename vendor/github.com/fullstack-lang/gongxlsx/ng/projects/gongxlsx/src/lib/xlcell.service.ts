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

import { XLCellDB } from './xlcell-db';

// insertion point for imports
import { XLRowDB } from './xlrow-db'
import { XLSheetDB } from './xlsheet-db'

@Injectable({
  providedIn: 'root'
})
export class XLCellService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  XLCellServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private xlcellsUrl: string

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
    this.xlcellsUrl = origin + '/api/github.com/fullstack-lang/gongxlsx/go/v1/xlcells';
  }

  /** GET xlcells from the server */
  getXLCells(): Observable<XLCellDB[]> {
    return this.http.get<XLCellDB[]>(this.xlcellsUrl)
      .pipe(
        tap(_ => this.log('fetched xlcells')),
        catchError(this.handleError<XLCellDB[]>('getXLCells', []))
      );
  }

  /** GET xlcell by id. Will 404 if id not found */
  getXLCell(id: number): Observable<XLCellDB> {
    const url = `${this.xlcellsUrl}/${id}`;
    return this.http.get<XLCellDB>(url).pipe(
      tap(_ => this.log(`fetched xlcell id=${id}`)),
      catchError(this.handleError<XLCellDB>(`getXLCell id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new xlcell to the server */
  postXLCell(xlcelldb: XLCellDB): Observable<XLCellDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _XLRow_Cells_reverse = xlcelldb.XLRow_Cells_reverse
    xlcelldb.XLRow_Cells_reverse = new XLRowDB
    let _XLSheet_SheetCells_reverse = xlcelldb.XLSheet_SheetCells_reverse
    xlcelldb.XLSheet_SheetCells_reverse = new XLSheetDB

    return this.http.post<XLCellDB>(this.xlcellsUrl, xlcelldb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        xlcelldb.XLRow_Cells_reverse = _XLRow_Cells_reverse
        xlcelldb.XLSheet_SheetCells_reverse = _XLSheet_SheetCells_reverse
        this.log(`posted xlcelldb id=${xlcelldb.ID}`)
      }),
      catchError(this.handleError<XLCellDB>('postXLCell'))
    );
  }

  /** DELETE: delete the xlcelldb from the server */
  deleteXLCell(xlcelldb: XLCellDB | number): Observable<XLCellDB> {
    const id = typeof xlcelldb === 'number' ? xlcelldb : xlcelldb.ID;
    const url = `${this.xlcellsUrl}/${id}`;

    return this.http.delete<XLCellDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted xlcelldb id=${id}`)),
      catchError(this.handleError<XLCellDB>('deleteXLCell'))
    );
  }

  /** PUT: update the xlcelldb on the server */
  updateXLCell(xlcelldb: XLCellDB): Observable<XLCellDB> {
    const id = typeof xlcelldb === 'number' ? xlcelldb : xlcelldb.ID;
    const url = `${this.xlcellsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _XLRow_Cells_reverse = xlcelldb.XLRow_Cells_reverse
    xlcelldb.XLRow_Cells_reverse = new XLRowDB
    let _XLSheet_SheetCells_reverse = xlcelldb.XLSheet_SheetCells_reverse
    xlcelldb.XLSheet_SheetCells_reverse = new XLSheetDB

    return this.http.put<XLCellDB>(url, xlcelldb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        xlcelldb.XLRow_Cells_reverse = _XLRow_Cells_reverse
        xlcelldb.XLSheet_SheetCells_reverse = _XLSheet_SheetCells_reverse
        this.log(`updated xlcelldb id=${xlcelldb.ID}`)
      }),
      catchError(this.handleError<XLCellDB>('updateXLCell'))
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